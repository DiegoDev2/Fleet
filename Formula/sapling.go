package main

// Sapling - Source control client
// Homepage: https://sapling-scm.com

import (
	"fmt"
	
	"os/exec"
)

func installSapling() {
	// Método 1: Descargar y extraer .tar.gz
	sapling_tar_url := "https://github.com/facebook/sapling/archive/refs/tags/0.2.20240718-145624+f4e9df48.tar.gz"
	sapling_cmd_tar := exec.Command("curl", "-L", sapling_tar_url, "-o", "package.tar.gz")
	err := sapling_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sapling_zip_url := "https://github.com/facebook/sapling/archive/refs/tags/0.2.20240718-145624+f4e9df48.zip"
	sapling_cmd_zip := exec.Command("curl", "-L", sapling_zip_url, "-o", "package.zip")
	err = sapling_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sapling_bin_url := "https://github.com/facebook/sapling/archive/refs/tags/0.2.20240718-145624+f4e9df48.bin"
	sapling_cmd_bin := exec.Command("curl", "-L", sapling_bin_url, "-o", "binary.bin")
	err = sapling_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sapling_src_url := "https://github.com/facebook/sapling/archive/refs/tags/0.2.20240718-145624+f4e9df48.src.tar.gz"
	sapling_cmd_src := exec.Command("curl", "-L", sapling_src_url, "-o", "source.tar.gz")
	err = sapling_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sapling_cmd_direct := exec.Command("./binary")
	err = sapling_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: gh")
	exec.Command("latte", "install", "gh").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
