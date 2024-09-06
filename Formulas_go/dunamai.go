package main

// Dunamai - Dynamic version generation
// Homepage: https://github.com/mtkennerly/dunamai

import (
	"fmt"
	
	"os/exec"
)

func installDunamai() {
	// Método 1: Descargar y extraer .tar.gz
	dunamai_tar_url := "https://files.pythonhosted.org/packages/a0/fe/aee602f08765de4dd753d2e5d6cbd480857182e345f161f7a19ad1979e4d/dunamai-1.22.0.tar.gz"
	dunamai_cmd_tar := exec.Command("curl", "-L", dunamai_tar_url, "-o", "package.tar.gz")
	err := dunamai_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dunamai_zip_url := "https://files.pythonhosted.org/packages/a0/fe/aee602f08765de4dd753d2e5d6cbd480857182e345f161f7a19ad1979e4d/dunamai-1.22.0.zip"
	dunamai_cmd_zip := exec.Command("curl", "-L", dunamai_zip_url, "-o", "package.zip")
	err = dunamai_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dunamai_bin_url := "https://files.pythonhosted.org/packages/a0/fe/aee602f08765de4dd753d2e5d6cbd480857182e345f161f7a19ad1979e4d/dunamai-1.22.0.bin"
	dunamai_cmd_bin := exec.Command("curl", "-L", dunamai_bin_url, "-o", "binary.bin")
	err = dunamai_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dunamai_src_url := "https://files.pythonhosted.org/packages/a0/fe/aee602f08765de4dd753d2e5d6cbd480857182e345f161f7a19ad1979e4d/dunamai-1.22.0.src.tar.gz"
	dunamai_cmd_src := exec.Command("curl", "-L", dunamai_src_url, "-o", "source.tar.gz")
	err = dunamai_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dunamai_cmd_direct := exec.Command("./binary")
	err = dunamai_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
