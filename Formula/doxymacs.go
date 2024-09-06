package main

// Doxymacs - Elisp package for using doxygen under Emacs
// Homepage: https://doxymacs.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDoxymacs() {
	// Método 1: Descargar y extraer .tar.gz
	doxymacs_tar_url := "https://downloads.sourceforge.net/project/doxymacs/doxymacs/1.8.0/doxymacs-1.8.0.tar.gz"
	doxymacs_cmd_tar := exec.Command("curl", "-L", doxymacs_tar_url, "-o", "package.tar.gz")
	err := doxymacs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doxymacs_zip_url := "https://downloads.sourceforge.net/project/doxymacs/doxymacs/1.8.0/doxymacs-1.8.0.zip"
	doxymacs_cmd_zip := exec.Command("curl", "-L", doxymacs_zip_url, "-o", "package.zip")
	err = doxymacs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doxymacs_bin_url := "https://downloads.sourceforge.net/project/doxymacs/doxymacs/1.8.0/doxymacs-1.8.0.bin"
	doxymacs_cmd_bin := exec.Command("curl", "-L", doxymacs_bin_url, "-o", "binary.bin")
	err = doxymacs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doxymacs_src_url := "https://downloads.sourceforge.net/project/doxymacs/doxymacs/1.8.0/doxymacs-1.8.0.src.tar.gz"
	doxymacs_cmd_src := exec.Command("curl", "-L", doxymacs_src_url, "-o", "source.tar.gz")
	err = doxymacs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doxymacs_cmd_direct := exec.Command("./binary")
	err = doxymacs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: doxygen")
	exec.Command("latte", "install", "doxygen").Run()
	fmt.Println("Instalando dependencia: emacs")
	exec.Command("latte", "install", "emacs").Run()
}
