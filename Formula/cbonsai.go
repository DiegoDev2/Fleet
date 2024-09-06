package main

// Cbonsai - Console Bonsai is a bonsai tree generator, written in C using ncurses
// Homepage: https://gitlab.com/jallbrit/cbonsai

import (
	"fmt"
	
	"os/exec"
)

func installCbonsai() {
	// Método 1: Descargar y extraer .tar.gz
	cbonsai_tar_url := "https://gitlab.com/jallbrit/cbonsai/-/archive/v1.3.1/cbonsai-v1.3.1.tar.gz"
	cbonsai_cmd_tar := exec.Command("curl", "-L", cbonsai_tar_url, "-o", "package.tar.gz")
	err := cbonsai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cbonsai_zip_url := "https://gitlab.com/jallbrit/cbonsai/-/archive/v1.3.1/cbonsai-v1.3.1.zip"
	cbonsai_cmd_zip := exec.Command("curl", "-L", cbonsai_zip_url, "-o", "package.zip")
	err = cbonsai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cbonsai_bin_url := "https://gitlab.com/jallbrit/cbonsai/-/archive/v1.3.1/cbonsai-v1.3.1.bin"
	cbonsai_cmd_bin := exec.Command("curl", "-L", cbonsai_bin_url, "-o", "binary.bin")
	err = cbonsai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cbonsai_src_url := "https://gitlab.com/jallbrit/cbonsai/-/archive/v1.3.1/cbonsai-v1.3.1.src.tar.gz"
	cbonsai_cmd_src := exec.Command("curl", "-L", cbonsai_src_url, "-o", "source.tar.gz")
	err = cbonsai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cbonsai_cmd_direct := exec.Command("./binary")
	err = cbonsai_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: scdoc")
	exec.Command("latte", "install", "scdoc").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
}
