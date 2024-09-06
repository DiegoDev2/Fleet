package main

// Ctags - Reimplementation of ctags(1)
// Homepage: https://ctags.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCtags() {
	// Método 1: Descargar y extraer .tar.gz
	ctags_tar_url := "https://downloads.sourceforge.net/project/ctags/ctags/5.8/ctags-5.8.tar.gz"
	ctags_cmd_tar := exec.Command("curl", "-L", ctags_tar_url, "-o", "package.tar.gz")
	err := ctags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ctags_zip_url := "https://downloads.sourceforge.net/project/ctags/ctags/5.8/ctags-5.8.zip"
	ctags_cmd_zip := exec.Command("curl", "-L", ctags_zip_url, "-o", "package.zip")
	err = ctags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ctags_bin_url := "https://downloads.sourceforge.net/project/ctags/ctags/5.8/ctags-5.8.bin"
	ctags_cmd_bin := exec.Command("curl", "-L", ctags_bin_url, "-o", "binary.bin")
	err = ctags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ctags_src_url := "https://downloads.sourceforge.net/project/ctags/ctags/5.8/ctags-5.8.src.tar.gz"
	ctags_cmd_src := exec.Command("curl", "-L", ctags_src_url, "-o", "source.tar.gz")
	err = ctags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ctags_cmd_direct := exec.Command("./binary")
	err = ctags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
}
