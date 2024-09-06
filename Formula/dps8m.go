package main

// Dps8m - Simulator of the 36-bit GE/Honeywell/Bull 600/6000-series mainframe computers
// Homepage: https://dps8m.gitlab.io/

import (
	"fmt"
	
	"os/exec"
)

func installDps8m() {
	// Método 1: Descargar y extraer .tar.gz
	dps8m_tar_url := "https://dps8m.gitlab.io/dps8m-r3.0.1-archive/R3.0.1/dps8m-r3.0.1-src.tar.gz"
	dps8m_cmd_tar := exec.Command("curl", "-L", dps8m_tar_url, "-o", "package.tar.gz")
	err := dps8m_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dps8m_zip_url := "https://dps8m.gitlab.io/dps8m-r3.0.1-archive/R3.0.1/dps8m-r3.0.1-src.zip"
	dps8m_cmd_zip := exec.Command("curl", "-L", dps8m_zip_url, "-o", "package.zip")
	err = dps8m_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dps8m_bin_url := "https://dps8m.gitlab.io/dps8m-r3.0.1-archive/R3.0.1/dps8m-r3.0.1-src.bin"
	dps8m_cmd_bin := exec.Command("curl", "-L", dps8m_bin_url, "-o", "binary.bin")
	err = dps8m_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dps8m_src_url := "https://dps8m.gitlab.io/dps8m-r3.0.1-archive/R3.0.1/dps8m-r3.0.1-src.src.tar.gz"
	dps8m_cmd_src := exec.Command("curl", "-L", dps8m_src_url, "-o", "source.tar.gz")
	err = dps8m_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dps8m_cmd_direct := exec.Command("./binary")
	err = dps8m_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
}
