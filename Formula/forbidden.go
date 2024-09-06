package main

// Forbidden - Bypass 4xx HTTP response status codes and more
// Homepage: https://github.com/ivan-sincek/forbidden

import (
	"fmt"
	
	"os/exec"
)

func installForbidden() {
	// Método 1: Descargar y extraer .tar.gz
	forbidden_tar_url := "https://files.pythonhosted.org/packages/fa/03/9f18651dbe09f130e444b836c448b3b3b8a6ddec4996c4183c17e9131592/forbidden-11.2.tar.gz"
	forbidden_cmd_tar := exec.Command("curl", "-L", forbidden_tar_url, "-o", "package.tar.gz")
	err := forbidden_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forbidden_zip_url := "https://files.pythonhosted.org/packages/fa/03/9f18651dbe09f130e444b836c448b3b3b8a6ddec4996c4183c17e9131592/forbidden-11.2.zip"
	forbidden_cmd_zip := exec.Command("curl", "-L", forbidden_zip_url, "-o", "package.zip")
	err = forbidden_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forbidden_bin_url := "https://files.pythonhosted.org/packages/fa/03/9f18651dbe09f130e444b836c448b3b3b8a6ddec4996c4183c17e9131592/forbidden-11.2.bin"
	forbidden_cmd_bin := exec.Command("curl", "-L", forbidden_bin_url, "-o", "binary.bin")
	err = forbidden_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forbidden_src_url := "https://files.pythonhosted.org/packages/fa/03/9f18651dbe09f130e444b836c448b3b3b8a6ddec4996c4183c17e9131592/forbidden-11.2.src.tar.gz"
	forbidden_cmd_src := exec.Command("curl", "-L", forbidden_src_url, "-o", "source.tar.gz")
	err = forbidden_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forbidden_cmd_direct := exec.Command("./binary")
	err = forbidden_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
