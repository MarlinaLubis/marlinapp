package url

import (
	"github.com/MarlinaLubis/marlinapp/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func Web(page *fiber.App) {
	page.Post("/api/whatsauth/request", controller.PostWhatsAuthRequest)  //API from user whatsapp message from iteung gowa
	page.Get("/ws/whatsauth/qr", websocket.New(controller.WsWhatsAuthQR)) //websocket whatsauth
	page.Get("/", controller.Home)                                        //ujicoba panggil package musik
	page.Get("/presensi", controller.GetPresensi)
	page.Get("/mahasiswa", controller.GetMahasiswa)
	page.Get("/dosen", controller.GetDosen)
	page.Get("/jamsidang", controller.GetJamSidang)
	page.Get("/bap", controller.GetBap)
	page.Get("/all", controller.GetAll)
	page.Get("/inspresensi", controller.InsertPresensi)
	page.Get("/insmahasiswa", controller.InsertMahasiswa)
	page.Get("/insdosen", controller.InsertDosen)
	page.Get("/insjamsidang", controller.InsertJamSidang)
	page.Get("/insbap", controller.InsertBap)
	page.Get("/presensi1", controller.GetAllPresensiin) //menampilkan seluruh data presensi
    page.Get("/presensi1/:id", controller.GetPresensiID) //menampilkan data presensi berdasarkan id
}
