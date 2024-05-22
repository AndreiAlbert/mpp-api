package vaccination

import "gorm.io/gorm"

type VaccinationRepository struct {
	db *gorm.DB
}

func NewVaccinationRepository(db *gorm.DB) *VaccinationRepository {
	return &VaccinationRepository{
		db: db,
	}
}

func (r *VaccinationRepository) CreateVaccination(vacc Vaccination) (Vaccination, error) {
	if err := r.db.Create(&vacc).Error; err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (r *VaccinationRepository) GetVaccinationById(id uint64) (Vaccination, error) {
	var vacc Vaccination
	if err := r.db.First(&vacc).Error; err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (r *VaccinationRepository) GetAllVaccinations() ([]Vaccination, error) {
	var vaccinations []Vaccination
	err := r.db.Find(&vaccinations).Error
	return vaccinations, err
}

func (r *VaccinationRepository) DeleteVaccination(id uint64) (Vaccination, error) {
	var vacc Vaccination
	if err := r.db.First(&vacc, "ID = ?", id).Error; err != nil {
		return Vaccination{}, err
	}
	if err := r.db.Delete(&vacc).Error; err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}

func (r *VaccinationRepository) GetVaccinationByCatId(catID uint) ([]Vaccination, error) {
	var vaccs []Vaccination
	if err := r.db.Where("pet_id = ?", catID).Find(&vaccs).Error; err != nil {
		return nil, err
	}
	return vaccs, nil
}

func (r *VaccinationRepository) UpdateVaccination(vaccId uint64, updatedVacc Vaccination) (Vaccination, error) {
	var vacc Vaccination
	if err := r.db.First(&vacc, "ID = ?", vaccId).Error; err != nil {
		return Vaccination{}, err
	}
	if err := r.db.Model(&vacc).Updates(updatedVacc).Error; err != nil {
		return Vaccination{}, err
	}
	return vacc, nil
}
