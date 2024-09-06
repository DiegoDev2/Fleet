package main

// Renameutils - Tools for file renaming
// Homepage: https://www.nongnu.org/renameutils/

import (
	"fmt"
	
	"os/exec"
)

func installRenameutils() {
	// Método 1: Descargar y extraer .tar.gz
	renameutils_tar_url := "https://download.savannah.gnu.org/releases/renameutils/renameutils-0.12.0.tar.gz"
	renameutils_cmd_tar := exec.Command("curl", "-L", renameutils_tar_url, "-o", "package.tar.gz")
	err := renameutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	renameutils_zip_url := "https://download.savannah.gnu.org/releases/renameutils/renameutils-0.12.0.zip"
	renameutils_cmd_zip := exec.Command("curl", "-L", renameutils_zip_url, "-o", "package.zip")
	err = renameutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	renameutils_bin_url := "https://download.savannah.gnu.org/releases/renameutils/renameutils-0.12.0.bin"
	renameutils_cmd_bin := exec.Command("curl", "-L", renameutils_bin_url, "-o", "binary.bin")
	err = renameutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	renameutils_src_url := "https://download.savannah.gnu.org/releases/renameutils/renameutils-0.12.0.src.tar.gz"
	renameutils_cmd_src := exec.Command("curl", "-L", renameutils_src_url, "-o", "source.tar.gz")
	err = renameutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	renameutils_cmd_direct := exec.Command("./binary")
	err = renameutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
