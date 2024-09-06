package main

// Libdicom - DICOM WSI read library
// Homepage: https://github.com/ImagingDataCommons/libdicom

import (
	"fmt"
	
	"os/exec"
)

func installLibdicom() {
	// Método 1: Descargar y extraer .tar.gz
	libdicom_tar_url := "https://github.com/ImagingDataCommons/libdicom/releases/download/v1.1.0/libdicom-1.1.0.tar.xz"
	libdicom_cmd_tar := exec.Command("curl", "-L", libdicom_tar_url, "-o", "package.tar.gz")
	err := libdicom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdicom_zip_url := "https://github.com/ImagingDataCommons/libdicom/releases/download/v1.1.0/libdicom-1.1.0.tar.xz"
	libdicom_cmd_zip := exec.Command("curl", "-L", libdicom_zip_url, "-o", "package.zip")
	err = libdicom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdicom_bin_url := "https://github.com/ImagingDataCommons/libdicom/releases/download/v1.1.0/libdicom-1.1.0.tar.xz"
	libdicom_cmd_bin := exec.Command("curl", "-L", libdicom_bin_url, "-o", "binary.bin")
	err = libdicom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdicom_src_url := "https://github.com/ImagingDataCommons/libdicom/releases/download/v1.1.0/libdicom-1.1.0.tar.xz"
	libdicom_cmd_src := exec.Command("curl", "-L", libdicom_src_url, "-o", "source.tar.gz")
	err = libdicom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdicom_cmd_direct := exec.Command("./binary")
	err = libdicom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: uthash")
	exec.Command("latte", "install", "uthash").Run()
}
