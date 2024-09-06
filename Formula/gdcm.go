package main

// Gdcm - Grassroots DICOM library and utilities for medical files
// Homepage: https://sourceforge.net/projects/gdcm/

import (
	"fmt"
	
	"os/exec"
)

func installGdcm() {
	// Método 1: Descargar y extraer .tar.gz
	gdcm_tar_url := "https://github.com/malaterre/GDCM/archive/refs/tags/v3.0.24.tar.gz"
	gdcm_cmd_tar := exec.Command("curl", "-L", gdcm_tar_url, "-o", "package.tar.gz")
	err := gdcm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gdcm_zip_url := "https://github.com/malaterre/GDCM/archive/refs/tags/v3.0.24.zip"
	gdcm_cmd_zip := exec.Command("curl", "-L", gdcm_zip_url, "-o", "package.zip")
	err = gdcm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gdcm_bin_url := "https://github.com/malaterre/GDCM/archive/refs/tags/v3.0.24.bin"
	gdcm_cmd_bin := exec.Command("curl", "-L", gdcm_bin_url, "-o", "binary.bin")
	err = gdcm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gdcm_src_url := "https://github.com/malaterre/GDCM/archive/refs/tags/v3.0.24.src.tar.gz"
	gdcm_cmd_src := exec.Command("curl", "-L", gdcm_src_url, "-o", "source.tar.gz")
	err = gdcm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gdcm_cmd_direct := exec.Command("./binary")
	err = gdcm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
