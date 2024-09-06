package main

// PythonCharsetNormalizer - Real First Universal Charset Detector, maintained alternative to Chardet
// Homepage: https://charset-normalizer.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installPythonCharsetNormalizer() {
	// Método 1: Descargar y extraer .tar.gz
	pythoncharsetnormalizer_tar_url := "https://files.pythonhosted.org/packages/63/09/c1bc53dab74b1816a00d8d030de5bf98f724c52c1635e07681d312f20be8/charset-normalizer-3.3.2.tar.gz"
	pythoncharsetnormalizer_cmd_tar := exec.Command("curl", "-L", pythoncharsetnormalizer_tar_url, "-o", "package.tar.gz")
	err := pythoncharsetnormalizer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythoncharsetnormalizer_zip_url := "https://files.pythonhosted.org/packages/63/09/c1bc53dab74b1816a00d8d030de5bf98f724c52c1635e07681d312f20be8/charset-normalizer-3.3.2.zip"
	pythoncharsetnormalizer_cmd_zip := exec.Command("curl", "-L", pythoncharsetnormalizer_zip_url, "-o", "package.zip")
	err = pythoncharsetnormalizer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythoncharsetnormalizer_bin_url := "https://files.pythonhosted.org/packages/63/09/c1bc53dab74b1816a00d8d030de5bf98f724c52c1635e07681d312f20be8/charset-normalizer-3.3.2.bin"
	pythoncharsetnormalizer_cmd_bin := exec.Command("curl", "-L", pythoncharsetnormalizer_bin_url, "-o", "binary.bin")
	err = pythoncharsetnormalizer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythoncharsetnormalizer_src_url := "https://files.pythonhosted.org/packages/63/09/c1bc53dab74b1816a00d8d030de5bf98f724c52c1635e07681d312f20be8/charset-normalizer-3.3.2.src.tar.gz"
	pythoncharsetnormalizer_cmd_src := exec.Command("curl", "-L", pythoncharsetnormalizer_src_url, "-o", "source.tar.gz")
	err = pythoncharsetnormalizer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythoncharsetnormalizer_cmd_direct := exec.Command("./binary")
	err = pythoncharsetnormalizer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
