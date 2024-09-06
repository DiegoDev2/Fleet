package main

// PythonPytz - Python library for cross platform timezone
// Homepage: https://pythonhosted.org/pytz/

import (
	"fmt"
	
	"os/exec"
)

func installPythonPytz() {
	// Método 1: Descargar y extraer .tar.gz
	pythonpytz_tar_url := "https://files.pythonhosted.org/packages/90/26/9f1f00a5d021fff16dee3de13d43e5e978f3d58928e129c3a62cf7eb9738/pytz-2024.1.tar.gz"
	pythonpytz_cmd_tar := exec.Command("curl", "-L", pythonpytz_tar_url, "-o", "package.tar.gz")
	err := pythonpytz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonpytz_zip_url := "https://files.pythonhosted.org/packages/90/26/9f1f00a5d021fff16dee3de13d43e5e978f3d58928e129c3a62cf7eb9738/pytz-2024.1.zip"
	pythonpytz_cmd_zip := exec.Command("curl", "-L", pythonpytz_zip_url, "-o", "package.zip")
	err = pythonpytz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonpytz_bin_url := "https://files.pythonhosted.org/packages/90/26/9f1f00a5d021fff16dee3de13d43e5e978f3d58928e129c3a62cf7eb9738/pytz-2024.1.bin"
	pythonpytz_cmd_bin := exec.Command("curl", "-L", pythonpytz_bin_url, "-o", "binary.bin")
	err = pythonpytz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonpytz_src_url := "https://files.pythonhosted.org/packages/90/26/9f1f00a5d021fff16dee3de13d43e5e978f3d58928e129c3a62cf7eb9738/pytz-2024.1.src.tar.gz"
	pythonpytz_cmd_src := exec.Command("curl", "-L", pythonpytz_src_url, "-o", "source.tar.gz")
	err = pythonpytz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonpytz_cmd_direct := exec.Command("./binary")
	err = pythonpytz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
