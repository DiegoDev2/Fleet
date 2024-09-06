package main

// Otf2 - Open Trace Format 2 file handling library
// Homepage: https://www.vi-hps.org/projects/score-p/

import (
	"fmt"
	
	"os/exec"
)

func installOtf2() {
	// Método 1: Descargar y extraer .tar.gz
	otf2_tar_url := "https://perftools.pages.jsc.fz-juelich.de/cicd/otf2/tags/otf2-3.0.3/otf2-3.0.3.tar.gz"
	otf2_cmd_tar := exec.Command("curl", "-L", otf2_tar_url, "-o", "package.tar.gz")
	err := otf2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	otf2_zip_url := "https://perftools.pages.jsc.fz-juelich.de/cicd/otf2/tags/otf2-3.0.3/otf2-3.0.3.zip"
	otf2_cmd_zip := exec.Command("curl", "-L", otf2_zip_url, "-o", "package.zip")
	err = otf2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	otf2_bin_url := "https://perftools.pages.jsc.fz-juelich.de/cicd/otf2/tags/otf2-3.0.3/otf2-3.0.3.bin"
	otf2_cmd_bin := exec.Command("curl", "-L", otf2_bin_url, "-o", "binary.bin")
	err = otf2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	otf2_src_url := "https://perftools.pages.jsc.fz-juelich.de/cicd/otf2/tags/otf2-3.0.3/otf2-3.0.3.src.tar.gz"
	otf2_cmd_src := exec.Command("curl", "-L", otf2_src_url, "-o", "source.tar.gz")
	err = otf2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	otf2_cmd_direct := exec.Command("./binary")
	err = otf2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: open-mpi")
exec.Command("latte", "install", "open-mpi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
