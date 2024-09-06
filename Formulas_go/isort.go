package main

// Isort - Sort Python imports automatically
// Homepage: https://pycqa.github.io/isort/

import (
	"fmt"
	
	"os/exec"
)

func installIsort() {
	// Método 1: Descargar y extraer .tar.gz
	isort_tar_url := "https://files.pythonhosted.org/packages/87/f9/c1eb8635a24e87ade2efce21e3ce8cd6b8630bb685ddc9cdaca1349b2eb5/isort-5.13.2.tar.gz"
	isort_cmd_tar := exec.Command("curl", "-L", isort_tar_url, "-o", "package.tar.gz")
	err := isort_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	isort_zip_url := "https://files.pythonhosted.org/packages/87/f9/c1eb8635a24e87ade2efce21e3ce8cd6b8630bb685ddc9cdaca1349b2eb5/isort-5.13.2.zip"
	isort_cmd_zip := exec.Command("curl", "-L", isort_zip_url, "-o", "package.zip")
	err = isort_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	isort_bin_url := "https://files.pythonhosted.org/packages/87/f9/c1eb8635a24e87ade2efce21e3ce8cd6b8630bb685ddc9cdaca1349b2eb5/isort-5.13.2.bin"
	isort_cmd_bin := exec.Command("curl", "-L", isort_bin_url, "-o", "binary.bin")
	err = isort_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	isort_src_url := "https://files.pythonhosted.org/packages/87/f9/c1eb8635a24e87ade2efce21e3ce8cd6b8630bb685ddc9cdaca1349b2eb5/isort-5.13.2.src.tar.gz"
	isort_cmd_src := exec.Command("curl", "-L", isort_src_url, "-o", "source.tar.gz")
	err = isort_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	isort_cmd_direct := exec.Command("./binary")
	err = isort_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
