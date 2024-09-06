package main

// Nvchecker - New version checker for software releases
// Homepage: https://github.com/lilydjwg/nvchecker

import (
	"fmt"
	
	"os/exec"
)

func installNvchecker() {
	// Método 1: Descargar y extraer .tar.gz
	nvchecker_tar_url := "https://files.pythonhosted.org/packages/7b/60/fd880c869c6a03768fcfe44168d7667f036e2499c8816dd106440e201332/nvchecker-2.15.1.tar.gz"
	nvchecker_cmd_tar := exec.Command("curl", "-L", nvchecker_tar_url, "-o", "package.tar.gz")
	err := nvchecker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvchecker_zip_url := "https://files.pythonhosted.org/packages/7b/60/fd880c869c6a03768fcfe44168d7667f036e2499c8816dd106440e201332/nvchecker-2.15.1.zip"
	nvchecker_cmd_zip := exec.Command("curl", "-L", nvchecker_zip_url, "-o", "package.zip")
	err = nvchecker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvchecker_bin_url := "https://files.pythonhosted.org/packages/7b/60/fd880c869c6a03768fcfe44168d7667f036e2499c8816dd106440e201332/nvchecker-2.15.1.bin"
	nvchecker_cmd_bin := exec.Command("curl", "-L", nvchecker_bin_url, "-o", "binary.bin")
	err = nvchecker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvchecker_src_url := "https://files.pythonhosted.org/packages/7b/60/fd880c869c6a03768fcfe44168d7667f036e2499c8816dd106440e201332/nvchecker-2.15.1.src.tar.gz"
	nvchecker_cmd_src := exec.Command("curl", "-L", nvchecker_src_url, "-o", "source.tar.gz")
	err = nvchecker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvchecker_cmd_direct := exec.Command("./binary")
	err = nvchecker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
