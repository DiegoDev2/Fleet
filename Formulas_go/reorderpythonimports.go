package main

// ReorderPythonImports - Rewrites source to reorder python imports
// Homepage: https://github.com/asottile/reorder-python-imports

import (
	"fmt"
	
	"os/exec"
)

func installReorderPythonImports() {
	// Método 1: Descargar y extraer .tar.gz
	reorderpythonimports_tar_url := "https://files.pythonhosted.org/packages/ae/f8/63ecf759c9149d7d7a8b612ebfe74901164dde9adcb1c40975ddc713db1c/reorder_python_imports-3.13.0.tar.gz"
	reorderpythonimports_cmd_tar := exec.Command("curl", "-L", reorderpythonimports_tar_url, "-o", "package.tar.gz")
	err := reorderpythonimports_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	reorderpythonimports_zip_url := "https://files.pythonhosted.org/packages/ae/f8/63ecf759c9149d7d7a8b612ebfe74901164dde9adcb1c40975ddc713db1c/reorder_python_imports-3.13.0.zip"
	reorderpythonimports_cmd_zip := exec.Command("curl", "-L", reorderpythonimports_zip_url, "-o", "package.zip")
	err = reorderpythonimports_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	reorderpythonimports_bin_url := "https://files.pythonhosted.org/packages/ae/f8/63ecf759c9149d7d7a8b612ebfe74901164dde9adcb1c40975ddc713db1c/reorder_python_imports-3.13.0.bin"
	reorderpythonimports_cmd_bin := exec.Command("curl", "-L", reorderpythonimports_bin_url, "-o", "binary.bin")
	err = reorderpythonimports_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	reorderpythonimports_src_url := "https://files.pythonhosted.org/packages/ae/f8/63ecf759c9149d7d7a8b612ebfe74901164dde9adcb1c40975ddc713db1c/reorder_python_imports-3.13.0.src.tar.gz"
	reorderpythonimports_cmd_src := exec.Command("curl", "-L", reorderpythonimports_src_url, "-o", "source.tar.gz")
	err = reorderpythonimports_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	reorderpythonimports_cmd_direct := exec.Command("./binary")
	err = reorderpythonimports_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
