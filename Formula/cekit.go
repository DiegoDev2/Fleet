package main

// Cekit - Container Evolution Kit
// Homepage: https://cekit.io

import (
	"fmt"
	
	"os/exec"
)

func installCekit() {
	// Método 1: Descargar y extraer .tar.gz
	cekit_tar_url := "https://files.pythonhosted.org/packages/45/68/5adda4ed0c9f5443110ac7b3c41f5492f09d34d939c39fccba6b0a5a00e9/cekit-4.13.0.tar.gz"
	cekit_cmd_tar := exec.Command("curl", "-L", cekit_tar_url, "-o", "package.tar.gz")
	err := cekit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cekit_zip_url := "https://files.pythonhosted.org/packages/45/68/5adda4ed0c9f5443110ac7b3c41f5492f09d34d939c39fccba6b0a5a00e9/cekit-4.13.0.zip"
	cekit_cmd_zip := exec.Command("curl", "-L", cekit_zip_url, "-o", "package.zip")
	err = cekit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cekit_bin_url := "https://files.pythonhosted.org/packages/45/68/5adda4ed0c9f5443110ac7b3c41f5492f09d34d939c39fccba6b0a5a00e9/cekit-4.13.0.bin"
	cekit_cmd_bin := exec.Command("curl", "-L", cekit_bin_url, "-o", "binary.bin")
	err = cekit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cekit_src_url := "https://files.pythonhosted.org/packages/45/68/5adda4ed0c9f5443110ac7b3c41f5492f09d34d939c39fccba6b0a5a00e9/cekit-4.13.0.src.tar.gz"
	cekit_cmd_src := exec.Command("curl", "-L", cekit_src_url, "-o", "source.tar.gz")
	err = cekit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cekit_cmd_direct := exec.Command("./binary")
	err = cekit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
