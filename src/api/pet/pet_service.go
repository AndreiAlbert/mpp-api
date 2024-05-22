package pet

import (
	"strconv"
)

type PetService struct {
	Repo *PetRepository
}

func NewPetService(repo *PetRepository) *PetService {
	return &PetService{
		Repo: repo,
	}
}

func (s *PetService) PetsByUserId(userId uint) ([]Pet, error) {
	pets, err := s.Repo.FindPetsByUserId(userId)
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (s *PetService) AllPets() ([]Pet, error) {
	pets, err := s.Repo.AllPets()
	if err != nil {
		return nil, err
	}
	return pets, nil
}

func (s *PetService) PetById(idStr string) (Pet, error) {
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Pet{}, err
	}
	pet, err := s.Repo.PetById(id)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (s *PetService) CreatePet(pet Pet) (Pet, error) {
	pet, err := s.Repo.CreatePet(pet)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (s *PetService) DeletePet(idStr string) (Pet, error) {
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Pet{}, err
	}
	pet, err := s.Repo.DeletePet(id)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (s *PetService) UpdatePet(idStr string, pet Pet) (Pet, error) {
	var id uint64
	id, err := strconv.ParseUint(idStr, 10, 0)
	if err != nil {
		return Pet{}, err
	}
	pet, err = s.Repo.UpdatePet(id, pet)
	if err != nil {
		return Pet{}, nil
	}
	return pet, err
}
