package main

// Vdirsyncer - Synchronize calendars and contacts
// Homepage: https://github.com/pimutils/vdirsyncer

import (
	"fmt"
	
	"os/exec"
)

func installVdirsyncer() {
	// Método 1: Descargar y extraer .tar.gz
	vdirsyncer_tar_url := "https://files.pythonhosted.org/packages/81/fb/6fbb7f1d102a59db275811a0de756d6f5bb55c624ba4bdf918b3fbd2ddc0/vdirsyncer-0.19.2.tar.gz"
	vdirsyncer_cmd_tar := exec.Command("curl", "-L", vdirsyncer_tar_url, "-o", "package.tar.gz")
	err := vdirsyncer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vdirsyncer_zip_url := "https://files.pythonhosted.org/packages/81/fb/6fbb7f1d102a59db275811a0de756d6f5bb55c624ba4bdf918b3fbd2ddc0/vdirsyncer-0.19.2.zip"
	vdirsyncer_cmd_zip := exec.Command("curl", "-L", vdirsyncer_zip_url, "-o", "package.zip")
	err = vdirsyncer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vdirsyncer_bin_url := "https://files.pythonhosted.org/packages/81/fb/6fbb7f1d102a59db275811a0de756d6f5bb55c624ba4bdf918b3fbd2ddc0/vdirsyncer-0.19.2.bin"
	vdirsyncer_cmd_bin := exec.Command("curl", "-L", vdirsyncer_bin_url, "-o", "binary.bin")
	err = vdirsyncer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vdirsyncer_src_url := "https://files.pythonhosted.org/packages/81/fb/6fbb7f1d102a59db275811a0de756d6f5bb55c624ba4bdf918b3fbd2ddc0/vdirsyncer-0.19.2.src.tar.gz"
	vdirsyncer_cmd_src := exec.Command("curl", "-L", vdirsyncer_src_url, "-o", "source.tar.gz")
	err = vdirsyncer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vdirsyncer_cmd_direct := exec.Command("./binary")
	err = vdirsyncer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
