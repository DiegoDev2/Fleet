package main

// Libpointing - Provides direct access to HID pointing devices
// Homepage: https://github.com/INRIA/libpointing

import (
	"fmt"
	
	"os/exec"
)

func installLibpointing() {
	// Método 1: Descargar y extraer .tar.gz
	libpointing_tar_url := "https://github.com/INRIA/libpointing/releases/download/v1.0.8/libpointing-mac-1.0.8.tar.gz"
	libpointing_cmd_tar := exec.Command("curl", "-L", libpointing_tar_url, "-o", "package.tar.gz")
	err := libpointing_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpointing_zip_url := "https://github.com/INRIA/libpointing/releases/download/v1.0.8/libpointing-mac-1.0.8.zip"
	libpointing_cmd_zip := exec.Command("curl", "-L", libpointing_zip_url, "-o", "package.zip")
	err = libpointing_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpointing_bin_url := "https://github.com/INRIA/libpointing/releases/download/v1.0.8/libpointing-mac-1.0.8.bin"
	libpointing_cmd_bin := exec.Command("curl", "-L", libpointing_bin_url, "-o", "binary.bin")
	err = libpointing_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpointing_src_url := "https://github.com/INRIA/libpointing/releases/download/v1.0.8/libpointing-mac-1.0.8.src.tar.gz"
	libpointing_cmd_src := exec.Command("curl", "-L", libpointing_src_url, "-o", "source.tar.gz")
	err = libpointing_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpointing_cmd_direct := exec.Command("./binary")
	err = libpointing_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
