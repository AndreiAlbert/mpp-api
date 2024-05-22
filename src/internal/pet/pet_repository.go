package pet

import "gorm.io/gorm"

type PetRepository struct {
	Db *gorm.DB
}

func NewPetRepository(db *gorm.DB) *PetRepository {
	return &PetRepository{
		Db: db,
	}
}

func (r *PetRepository) FindPetsByUserId(userId uint) ([]Pet, error) {
	var pets []Pet
	if err := r.Db.Where("user_id = ?", userId).Find(&pets).Error; err != nil {
		return nil, err
	}
	return pets, nil
}

func (r *PetRepository) AllPets() ([]Pet, error) {
	var pets []Pet
	result := r.Db.Find(&pets)
	return pets, result.Error
}

func (r *PetRepository) PetById(id uint64) (Pet, error) {
	var pet Pet
	if err := r.Db.First(&pet, "ID = ?", id).Error; err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) CreatePet(pet Pet) (Pet, error) {
	if err := r.Db.Create(&pet).Error; err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) DeletePet(id uint64) (Pet, error) {
	var pet Pet
	if err := r.Db.First(&pet, "ID = ?", id).Error; err != nil {
		return Pet{}, err
	}
	if err := r.Db.Delete(&pet).Error; err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (r *PetRepository) UpdatePet(id uint64, updatedPet Pet) (Pet, error) {
	var pet Pet
	if err := r.Db.First(&pet, "ID = ?", id).Error; err != nil {
		return Pet{}, err
	}
	if err := r.Db.Model(&pet).Updates(updatedPet).Error; err != nil {
		return Pet{}, err
	}
	return pet, nil
}
