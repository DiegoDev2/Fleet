package main

// Shtool - GNU's portable shell tool
// Homepage: https://www.gnu.org/software/shtool/

import (
	"fmt"
	
	"os/exec"
)

func installShtool() {
	// Método 1: Descargar y extraer .tar.gz
	shtool_tar_url := "https://ftp.gnu.org/gnu/shtool/shtool-2.0.8.tar.gz"
	shtool_cmd_tar := exec.Command("curl", "-L", shtool_tar_url, "-o", "package.tar.gz")
	err := shtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shtool_zip_url := "https://ftp.gnu.org/gnu/shtool/shtool-2.0.8.zip"
	shtool_cmd_zip := exec.Command("curl", "-L", shtool_zip_url, "-o", "package.zip")
	err = shtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shtool_bin_url := "https://ftp.gnu.org/gnu/shtool/shtool-2.0.8.bin"
	shtool_cmd_bin := exec.Command("curl", "-L", shtool_bin_url, "-o", "binary.bin")
	err = shtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shtool_src_url := "https://ftp.gnu.org/gnu/shtool/shtool-2.0.8.src.tar.gz"
	shtool_cmd_src := exec.Command("curl", "-L", shtool_src_url, "-o", "source.tar.gz")
	err = shtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shtool_cmd_direct := exec.Command("./binary")
	err = shtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
