package main

// Sn0int - Semi-automatic OSINT framework and package manager
// Homepage: https://github.com/kpcyrd/sn0int

import (
	"fmt"
	
	"os/exec"
)

func installSn0int() {
	// Método 1: Descargar y extraer .tar.gz
	sn0int_tar_url := "https://github.com/kpcyrd/sn0int/archive/refs/tags/v0.26.0.tar.gz"
	sn0int_cmd_tar := exec.Command("curl", "-L", sn0int_tar_url, "-o", "package.tar.gz")
	err := sn0int_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sn0int_zip_url := "https://github.com/kpcyrd/sn0int/archive/refs/tags/v0.26.0.zip"
	sn0int_cmd_zip := exec.Command("curl", "-L", sn0int_zip_url, "-o", "package.zip")
	err = sn0int_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sn0int_bin_url := "https://github.com/kpcyrd/sn0int/archive/refs/tags/v0.26.0.bin"
	sn0int_cmd_bin := exec.Command("curl", "-L", sn0int_bin_url, "-o", "binary.bin")
	err = sn0int_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sn0int_src_url := "https://github.com/kpcyrd/sn0int/archive/refs/tags/v0.26.0.src.tar.gz"
	sn0int_cmd_src := exec.Command("curl", "-L", sn0int_src_url, "-o", "source.tar.gz")
	err = sn0int_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sn0int_cmd_direct := exec.Command("./binary")
	err = sn0int_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libseccomp")
	exec.Command("latte", "install", "libseccomp").Run()
}
