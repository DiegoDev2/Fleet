package main

// EcflowUi - User interface for client/server workflow package
// Homepage: https://confluence.ecmwf.int/display/ECFLOW

import (
	"fmt"
	
	"os/exec"
)

func installEcflowUi() {
	// Método 1: Descargar y extraer .tar.gz
	ecflowui_tar_url := "https://confluence.ecmwf.int/download/attachments/8650755/ecFlow-5.13.4-Source.tar.gz"
	ecflowui_cmd_tar := exec.Command("curl", "-L", ecflowui_tar_url, "-o", "package.tar.gz")
	err := ecflowui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ecflowui_zip_url := "https://confluence.ecmwf.int/download/attachments/8650755/ecFlow-5.13.4-Source.zip"
	ecflowui_cmd_zip := exec.Command("curl", "-L", ecflowui_zip_url, "-o", "package.zip")
	err = ecflowui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ecflowui_bin_url := "https://confluence.ecmwf.int/download/attachments/8650755/ecFlow-5.13.4-Source.bin"
	ecflowui_cmd_bin := exec.Command("curl", "-L", ecflowui_bin_url, "-o", "binary.bin")
	err = ecflowui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ecflowui_src_url := "https://confluence.ecmwf.int/download/attachments/8650755/ecFlow-5.13.4-Source.src.tar.gz"
	ecflowui_cmd_src := exec.Command("curl", "-L", ecflowui_src_url, "-o", "source.tar.gz")
	err = ecflowui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ecflowui_cmd_direct := exec.Command("./binary")
	err = ecflowui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
}
