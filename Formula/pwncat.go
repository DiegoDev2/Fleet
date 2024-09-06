package main

// Pwncat - Netcat with FW/IDS/IPS evasion, self-inject-, bind- and reverse shell
// Homepage: https://pwncat.org

import (
	"fmt"
	
	"os/exec"
)

func installPwncat() {
	// Método 1: Descargar y extraer .tar.gz
	pwncat_tar_url := "https://files.pythonhosted.org/packages/c9/ce/51f7b53a8ee3b4afe4350577ee92f416f32b9b166f0d84b480fec1717a42/pwncat-0.1.2.tar.gz"
	pwncat_cmd_tar := exec.Command("curl", "-L", pwncat_tar_url, "-o", "package.tar.gz")
	err := pwncat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwncat_zip_url := "https://files.pythonhosted.org/packages/c9/ce/51f7b53a8ee3b4afe4350577ee92f416f32b9b166f0d84b480fec1717a42/pwncat-0.1.2.zip"
	pwncat_cmd_zip := exec.Command("curl", "-L", pwncat_zip_url, "-o", "package.zip")
	err = pwncat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwncat_bin_url := "https://files.pythonhosted.org/packages/c9/ce/51f7b53a8ee3b4afe4350577ee92f416f32b9b166f0d84b480fec1717a42/pwncat-0.1.2.bin"
	pwncat_cmd_bin := exec.Command("curl", "-L", pwncat_bin_url, "-o", "binary.bin")
	err = pwncat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwncat_src_url := "https://files.pythonhosted.org/packages/c9/ce/51f7b53a8ee3b4afe4350577ee92f416f32b9b166f0d84b480fec1717a42/pwncat-0.1.2.src.tar.gz"
	pwncat_cmd_src := exec.Command("curl", "-L", pwncat_src_url, "-o", "source.tar.gz")
	err = pwncat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwncat_cmd_direct := exec.Command("./binary")
	err = pwncat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
