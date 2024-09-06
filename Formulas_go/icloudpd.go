package main

// Icloudpd - Tool to download photos from iCloud
// Homepage: https://github.com/icloud-photos-downloader/icloud_photos_downloader

import (
	"fmt"
	
	"os/exec"
)

func installIcloudpd() {
	// Método 1: Descargar y extraer .tar.gz
	icloudpd_tar_url := "https://github.com/icloud-photos-downloader/icloud_photos_downloader.git"
	icloudpd_cmd_tar := exec.Command("curl", "-L", icloudpd_tar_url, "-o", "package.tar.gz")
	err := icloudpd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	icloudpd_zip_url := "https://github.com/icloud-photos-downloader/icloud_photos_downloader.git"
	icloudpd_cmd_zip := exec.Command("curl", "-L", icloudpd_zip_url, "-o", "package.zip")
	err = icloudpd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	icloudpd_bin_url := "https://github.com/icloud-photos-downloader/icloud_photos_downloader.git"
	icloudpd_cmd_bin := exec.Command("curl", "-L", icloudpd_bin_url, "-o", "binary.bin")
	err = icloudpd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	icloudpd_src_url := "https://github.com/icloud-photos-downloader/icloud_photos_downloader.git"
	icloudpd_cmd_src := exec.Command("curl", "-L", icloudpd_src_url, "-o", "source.tar.gz")
	err = icloudpd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	icloudpd_cmd_direct := exec.Command("./binary")
	err = icloudpd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
}
