package main

// Nsuds - Ncurses Sudoku system
// Homepage: https://nsuds.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installNsuds() {
	// Método 1: Descargar y extraer .tar.gz
	nsuds_tar_url := "https://downloads.sourceforge.net/project/nsuds/nsuds/nsuds-0.7B/nsuds-0.7B.tar.gz"
	nsuds_cmd_tar := exec.Command("curl", "-L", nsuds_tar_url, "-o", "package.tar.gz")
	err := nsuds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nsuds_zip_url := "https://downloads.sourceforge.net/project/nsuds/nsuds/nsuds-0.7B/nsuds-0.7B.zip"
	nsuds_cmd_zip := exec.Command("curl", "-L", nsuds_zip_url, "-o", "package.zip")
	err = nsuds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nsuds_bin_url := "https://downloads.sourceforge.net/project/nsuds/nsuds/nsuds-0.7B/nsuds-0.7B.bin"
	nsuds_cmd_bin := exec.Command("curl", "-L", nsuds_bin_url, "-o", "binary.bin")
	err = nsuds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nsuds_src_url := "https://downloads.sourceforge.net/project/nsuds/nsuds/nsuds-0.7B/nsuds-0.7B.src.tar.gz"
	nsuds_cmd_src := exec.Command("curl", "-L", nsuds_src_url, "-o", "source.tar.gz")
	err = nsuds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nsuds_cmd_direct := exec.Command("./binary")
	err = nsuds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
