package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"os"
	"patients/database/structs"
	"patients/logging"
	"sync"
)

type Data struct {
	patients   []structs.Patient
	configPath string
	mu         sync.RWMutex
}

func (d *Data) Load(byteArr []byte) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	err := json.Unmarshal(byteArr, &d.patients)
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) Save() error {
	d.mu.RLock()
	defer d.mu.RUnlock()

	marshaledData, err := json.Marshal(d.patients)
	if err != nil {
		return err
	}

	err = os.WriteFile("temp.json", marshaledData, 0744)
	if err != nil {
		return err
	}

	err = os.Remove(d.configPath)
	if err != nil {
		return err
	}

	err = os.Rename("temp.json", d.configPath)
	if err != nil {
		return err
	}

	return nil
}

func (d *Data) AddPatient(p structs.Patient) {
	d.mu.Lock()
	defer d.mu.Unlock()

	d.patients = append(d.patients, p)
}

func (d *Data) EditPatient(p structs.Patient) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i := range d.patients {
		if d.patients[i].GUID == p.GUID {
			d.patients[i] = p
			return nil
		}
	}

	return fmt.Errorf("patient does not exist")
}

func (d *Data) ListPatients() []structs.Patient {
	d.mu.RLock()
	defer d.mu.RUnlock()

	return d.patients
}

func (d *Data) DeletePatient(guid uuid.UUID) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	for i := range d.patients {
		if d.patients[i].GUID == guid {
			d.patients = append(d.patients[:i], d.patients[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("patient does not exist")
}

var (
	PatientsData Data
)

func Initialize(configPath string) error {
	file, err := os.ReadFile(configPath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		logging.LogError(err, "Failed read patient list from disk")
		return err
	} else if err != nil {
		if file, err = os.ReadFile("temp.json"); err != nil {
			logging.LogError(err, "Failed read patient list from disk")
			return err
		}
	}

	err = PatientsData.Load(file)
	if err != nil {
		logging.LogError(err, "Failed load patient list to struct")
		return err
	}

	PatientsData.configPath = configPath

	return nil
}
