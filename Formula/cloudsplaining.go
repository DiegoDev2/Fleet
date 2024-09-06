package main

// Cloudsplaining - AWS IAM Security Assessment tool
// Homepage: https://cloudsplaining.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installCloudsplaining() {
	// Método 1: Descargar y extraer .tar.gz
	cloudsplaining_tar_url := "https://files.pythonhosted.org/packages/84/0b/b232db103402bc27a9216909c6a9a3d308936e5895e5cadfb6783bccd719/cloudsplaining-0.6.3.tar.gz"
	cloudsplaining_cmd_tar := exec.Command("curl", "-L", cloudsplaining_tar_url, "-o", "package.tar.gz")
	err := cloudsplaining_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudsplaining_zip_url := "https://files.pythonhosted.org/packages/84/0b/b232db103402bc27a9216909c6a9a3d308936e5895e5cadfb6783bccd719/cloudsplaining-0.6.3.zip"
	cloudsplaining_cmd_zip := exec.Command("curl", "-L", cloudsplaining_zip_url, "-o", "package.zip")
	err = cloudsplaining_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudsplaining_bin_url := "https://files.pythonhosted.org/packages/84/0b/b232db103402bc27a9216909c6a9a3d308936e5895e5cadfb6783bccd719/cloudsplaining-0.6.3.bin"
	cloudsplaining_cmd_bin := exec.Command("curl", "-L", cloudsplaining_bin_url, "-o", "binary.bin")
	err = cloudsplaining_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudsplaining_src_url := "https://files.pythonhosted.org/packages/84/0b/b232db103402bc27a9216909c6a9a3d308936e5895e5cadfb6783bccd719/cloudsplaining-0.6.3.src.tar.gz"
	cloudsplaining_cmd_src := exec.Command("curl", "-L", cloudsplaining_src_url, "-o", "source.tar.gz")
	err = cloudsplaining_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudsplaining_cmd_direct := exec.Command("./binary")
	err = cloudsplaining_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
