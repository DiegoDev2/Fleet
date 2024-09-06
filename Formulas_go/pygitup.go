package main

// Pygitup - Nicer 'git pull'
// Homepage: https://github.com/msiemens/PyGitUp

import (
	"fmt"
	
	"os/exec"
)

func installPygitup() {
	// Método 1: Descargar y extraer .tar.gz
	pygitup_tar_url := "https://files.pythonhosted.org/packages/55/13/2dd3d4c9a021eb5fa6d8afbb29eb9e6eb57faa56cf10effe879c9626eed1/git_up-2.2.0.tar.gz"
	pygitup_cmd_tar := exec.Command("curl", "-L", pygitup_tar_url, "-o", "package.tar.gz")
	err := pygitup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pygitup_zip_url := "https://files.pythonhosted.org/packages/55/13/2dd3d4c9a021eb5fa6d8afbb29eb9e6eb57faa56cf10effe879c9626eed1/git_up-2.2.0.zip"
	pygitup_cmd_zip := exec.Command("curl", "-L", pygitup_zip_url, "-o", "package.zip")
	err = pygitup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pygitup_bin_url := "https://files.pythonhosted.org/packages/55/13/2dd3d4c9a021eb5fa6d8afbb29eb9e6eb57faa56cf10effe879c9626eed1/git_up-2.2.0.bin"
	pygitup_cmd_bin := exec.Command("curl", "-L", pygitup_bin_url, "-o", "binary.bin")
	err = pygitup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pygitup_src_url := "https://files.pythonhosted.org/packages/55/13/2dd3d4c9a021eb5fa6d8afbb29eb9e6eb57faa56cf10effe879c9626eed1/git_up-2.2.0.src.tar.gz"
	pygitup_cmd_src := exec.Command("curl", "-L", pygitup_src_url, "-o", "source.tar.gz")
	err = pygitup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pygitup_cmd_direct := exec.Command("./binary")
	err = pygitup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
