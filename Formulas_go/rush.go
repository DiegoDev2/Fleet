package main

// Rush - GNU's Restricted User SHell
// Homepage: https://www.gnu.org.ua/software/rush/

import (
	"fmt"
	
	"os/exec"
)

func installRush() {
	// Método 1: Descargar y extraer .tar.gz
	rush_tar_url := "https://ftp.gnu.org/gnu/rush/rush-2.4.tar.xz"
	rush_cmd_tar := exec.Command("curl", "-L", rush_tar_url, "-o", "package.tar.gz")
	err := rush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rush_zip_url := "https://ftp.gnu.org/gnu/rush/rush-2.4.tar.xz"
	rush_cmd_zip := exec.Command("curl", "-L", rush_zip_url, "-o", "package.zip")
	err = rush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rush_bin_url := "https://ftp.gnu.org/gnu/rush/rush-2.4.tar.xz"
	rush_cmd_bin := exec.Command("curl", "-L", rush_bin_url, "-o", "binary.bin")
	err = rush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rush_src_url := "https://ftp.gnu.org/gnu/rush/rush-2.4.tar.xz"
	rush_cmd_src := exec.Command("curl", "-L", rush_src_url, "-o", "source.tar.gz")
	err = rush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rush_cmd_direct := exec.Command("./binary")
	err = rush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
