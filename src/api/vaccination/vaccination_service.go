package vaccination

import "strconv"

type VaccinationService struct {
	Repo *VaccinationRepository
}

func NewVaccinationService(repo *VaccinationRepository) *VaccinationService {
	return &VaccinationService{
		Repo: repo,
	}
}

func (s *VaccinationService) VaccinationsByCatId(catId string) ([]Vaccination, error) {
	catIdNr, err := strconv.ParseUint(catId, 10, 0)
	if err != nil {
		return nil, err
	}
	return s.Repo.GetVaccinationByCatId(uint(catIdNr))
}

func (s *VaccinationService) AllVaccinations() ([]Vaccination, error) {
	return s.Repo.GetAllVaccinations()
}

func (s *VaccinationService) VaccinationById(idStr string) (Vaccination, error) {
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Vaccination{}, err
	}
	vacc, err := s.Repo.GetVaccinationById(id)
	if err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (s *VaccinationService) CreateVaccination(vacc Vaccination) (Vaccination, error) {
	vacc, err := s.Repo.CreateVaccination(vacc)
	if err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (s *VaccinationService) DeleteVaccination(idStr string) (Vaccination, error) {
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Vaccination{}, err
	}
	vacc, err := s.Repo.DeleteVaccination(id)
	if err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (s *VaccinationService) UpdateVaccination(idStr string, vacc Vaccination) (Vaccination, error) {
	var id uint64
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Vaccination{}, err
	}
	vacc, err = s.Repo.UpdateVaccination(id, vacc)
	if err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}
