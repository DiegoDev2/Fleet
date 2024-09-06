package main

// Qcli - Report audiovisual metrics via libavfilter
// Homepage: https://bavc.org/preserve-media/preservation-tools

import (
	"fmt"
	
	"os/exec"
)

func installQcli() {
	// Método 1: Descargar y extraer .tar.gz
	qcli_tar_url := "https://github.com/bavc/qctools.git"
	qcli_cmd_tar := exec.Command("curl", "-L", qcli_tar_url, "-o", "package.tar.gz")
	err := qcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qcli_zip_url := "https://github.com/bavc/qctools.git"
	qcli_cmd_zip := exec.Command("curl", "-L", qcli_zip_url, "-o", "package.zip")
	err = qcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qcli_bin_url := "https://github.com/bavc/qctools.git"
	qcli_cmd_bin := exec.Command("curl", "-L", qcli_bin_url, "-o", "binary.bin")
	err = qcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qcli_src_url := "https://github.com/bavc/qctools.git"
	qcli_cmd_src := exec.Command("curl", "-L", qcli_src_url, "-o", "source.tar.gz")
	err = qcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qcli_cmd_direct := exec.Command("./binary")
	err = qcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg@6")
	exec.Command("latte", "install", "ffmpeg@6").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: qwt")
	exec.Command("latte", "install", "qwt").Run()
}
