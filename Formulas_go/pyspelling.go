package main

// Pyspelling - Spell checker automation tool
// Homepage: https://facelessuser.github.io/pyspelling/

import (
	"fmt"
	
	"os/exec"
)

func installPyspelling() {
	// Método 1: Descargar y extraer .tar.gz
	pyspelling_tar_url := "https://files.pythonhosted.org/packages/12/07/168a857755a29b7e41550a28cd8f527025bc62fcb36a951d8f3f2eedcdf7/pyspelling-2.10.tar.gz"
	pyspelling_cmd_tar := exec.Command("curl", "-L", pyspelling_tar_url, "-o", "package.tar.gz")
	err := pyspelling_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pyspelling_zip_url := "https://files.pythonhosted.org/packages/12/07/168a857755a29b7e41550a28cd8f527025bc62fcb36a951d8f3f2eedcdf7/pyspelling-2.10.zip"
	pyspelling_cmd_zip := exec.Command("curl", "-L", pyspelling_zip_url, "-o", "package.zip")
	err = pyspelling_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pyspelling_bin_url := "https://files.pythonhosted.org/packages/12/07/168a857755a29b7e41550a28cd8f527025bc62fcb36a951d8f3f2eedcdf7/pyspelling-2.10.bin"
	pyspelling_cmd_bin := exec.Command("curl", "-L", pyspelling_bin_url, "-o", "binary.bin")
	err = pyspelling_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pyspelling_src_url := "https://files.pythonhosted.org/packages/12/07/168a857755a29b7e41550a28cd8f527025bc62fcb36a951d8f3f2eedcdf7/pyspelling-2.10.src.tar.gz"
	pyspelling_cmd_src := exec.Command("curl", "-L", pyspelling_src_url, "-o", "source.tar.gz")
	err = pyspelling_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pyspelling_cmd_direct := exec.Command("./binary")
	err = pyspelling_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: aspell")
exec.Command("latte", "install", "aspell")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
