package controller

import (
	"errors"
	"fmt"
	"net/http"

	inimodel "github.com/MarlinaLubis/be_bap/model"
	inimodul "github.com/MarlinaLubis/be_bap/module"
	"github.com/MarlinaLubis/marlinapp/config"
	"github.com/aiteung/musik"
	cek "github.com/aiteung/presensi"
	"github.com/gofiber/fiber/v2"
	inimodullatihan "github.com/indrariksa/be_presensi/module"
    itumodel "github.com/indrariksa/be_presensi/model"
	itumodul "github.com/indrariksa/be_presensi/module"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Home(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"github_repo": "https://github.com/MarlinaLubis/marlinapp",
		"message":     "You are at the root endpoint 😉",
		"success":     true,
	})
}

func Homepage(c *fiber.Ctx) error {
	ipaddr := musik.GetIPaddress()
	return c.JSON(ipaddr)
}

func GetPresensiold(c *fiber.Ctx) error {
	nl := cek.GetPresensiCurrentMonth(config.Ulbimongoconn)
	return c.JSON(nl)
}

func GetPresensi(c *fiber.Ctx) error {
	nl := inimodul.GetPresensiFromCheckin("Masuk", config.Ulbimongoconn, "presensi")
	return c.JSON(nl)
}

func GetMahasiswa(c *fiber.Ctx) error {
	nl := inimodul.GetMahasiswaFromNama("Angelina", config.Ulbimongoconn, "mahasiswa")
	return c.JSON(nl)
}

func GetDosen(c *fiber.Ctx) error {
	nl := inimodul.GetDosenFromNama("Indra Riksa", config.Ulbimongoconn, "dosen")
	return c.JSON(nl)
}

func GetJamSidang(c *fiber.Ctx) error {
	nl := inimodul.GetJamSidangFromDurasi(7, config.Ulbimongoconn, "jamsidang")
	return c.JSON(nl)
}

func GetBap(c *fiber.Ctx) error {
	nl := inimodul.GetBapFromJudul("Aplikasi Perpustakaan Berbasis Web", config.Ulbimongoconn, "bap")
	return c.JSON(nl)
}

func InsertPresensi(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi inimodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return err
	}
	insertedID := inimodul.InsertPresensi(db, "presensi",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Mahasiswa Berhasil Disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertMahasiswa(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var mahasiswa inimodel.Mahasiswa
	if err := c.BodyParser(&mahasiswa); err != nil {
		return err
	}
	insertedID := inimodul.InsertMahasiswa(db, "mahasiswa",
		mahasiswa.Nama,
		mahasiswa.Phone_number,
		mahasiswa.Email,
		mahasiswa.Jurusan,
		mahasiswa.Jam_sidang,
		mahasiswa.Hari_sidang)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Mahasiswa Berhasil Disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertDosen(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var dosen inimodel.Dosen
	if err := c.BodyParser(&dosen); err != nil {
		return err
	}
	insertedID := inimodul.InsertDosen(db, "dosen",

		dosen.Nama,
		dosen.Phone_number,
		dosen.Email,
		dosen.Jam_sidang,
		dosen.Hari_sidang)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Dosen Berhasil Disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertJamSidang(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var jamsidang inimodel.JamSidang
	if err := c.BodyParser(&jamsidang); err != nil {
		return err
	}
	insertedID := inimodul.InsertJamSidang(db, "jamsidang",
		jamsidang.Durasi,
		jamsidang.Jam_masuk,
		jamsidang.Jam_keluar,
		jamsidang.Gmt,
		jamsidang.Hari)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Jam Sidang Berhasil Disimpan.",
		"inserted_id": insertedID,
	})
}

func InsertBap(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var bap inimodel.Bap
	if err := c.BodyParser(&bap); err != nil {
		return err
	}
	insertedID := inimodul.InsertBap(db, "bap",
		bap.Judul,
		bap.Nama_Mahasiswa,
		bap.Prodi,
		bap.Hasil_revisi,
		bap.Dosen)
	return c.JSON(map[string]interface{}{
		"status":      http.StatusOK,
		"message":     "Data Bap Berhasil Disimpan.",
		"inserted_id": insertedID,
	})
}

func GetAll(c *fiber.Ctx) error {
	nl := inimodul.GetAllBapFromJudul("Aplikasi Perpustakaan Berbasis Web", config.Ulbimongoconn, "bap")
	return c.JSON(nl)
}

func GetAllPresensi(c *fiber.Ctx) error {
	ps := inimodullatihan.GetAllPresensi(config.Ulbimongoconn, "presensi1")
	return c.JSON(ps)
}

func GetPresensiID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}
	ps, err := inimodullatihan.GetPresensiFromID(objID, config.Ulbimongoconn, "presensi1")
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"status":  http.StatusNotFound,
				"message": fmt.Sprintf("No data found for id %s", id),
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error retrieving data for id %s", id),
		})
	}
	return c.JSON(ps)
}

func GetAllBap(c *fiber.Ctx) error {
	ps := inimodul.GetAllBap(config.Ulbimongoconn, "bap")
	return c.JSON(ps)
}

func InsertData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn
	var presensi itumodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	insertedID, err := itumodul.InsertPresensi(db, "presensi1",
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":      http.StatusOK,
		"message":     "Data berhasil disimpan.",
		"inserted_id": insertedID,
	})
}

func UpdateData(c *fiber.Ctx) error {
	db := config.Ulbimongoconn

	// Get the ID from the URL parameter
	id := c.Params("id")

	// Parse the ID into an ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Parse the request body into a Presensi object
	var presensi itumodel.Presensi
	if err := c.BodyParser(&presensi); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Call the UpdatePresensi function with the parsed ID and the Presensi object
	err = itumodul.UpdatePresensi(db, "presensi1",
		objectID,
		presensi.Longitude,
		presensi.Latitude,
		presensi.Location,
		presensi.Phone_number,
		presensi.Checkin,
		presensi.Biodata)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Data successfully updated",
	})
}

func DeletePresensiByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": "Wrong parameter",
		})
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":  http.StatusBadRequest,
			"message": "Invalid id parameter",
		})
	}

	err = itumodul.DeletePresensiByID(objID, config.Ulbimongoconn, "presensi1")
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"status":  http.StatusInternalServerError,
			"message": fmt.Sprintf("Error deleting data for id %s", id),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": fmt.Sprintf("Data with id %s deleted successfully", id),
	})
}