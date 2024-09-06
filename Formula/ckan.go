package main

// Ckan - Comprehensive Kerbal Archive Network
// Homepage: https://github.com/KSP-CKAN/CKAN/

import (
	"fmt"
	
	"os/exec"
)

func installCkan() {
	// Método 1: Descargar y extraer .tar.gz
	ckan_tar_url := "https://github.com/KSP-CKAN/CKAN/releases/download/v1.34.4/ckan.exe"
	ckan_cmd_tar := exec.Command("curl", "-L", ckan_tar_url, "-o", "package.tar.gz")
	err := ckan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ckan_zip_url := "https://github.com/KSP-CKAN/CKAN/releases/download/v1.34.4/ckan.exe"
	ckan_cmd_zip := exec.Command("curl", "-L", ckan_zip_url, "-o", "package.zip")
	err = ckan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ckan_bin_url := "https://github.com/KSP-CKAN/CKAN/releases/download/v1.34.4/ckan.exe"
	ckan_cmd_bin := exec.Command("curl", "-L", ckan_bin_url, "-o", "binary.bin")
	err = ckan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ckan_src_url := "https://github.com/KSP-CKAN/CKAN/releases/download/v1.34.4/ckan.exe"
	ckan_cmd_src := exec.Command("curl", "-L", ckan_src_url, "-o", "source.tar.gz")
	err = ckan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ckan_cmd_direct := exec.Command("./binary")
	err = ckan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mono")
	exec.Command("latte", "install", "mono").Run()
}
