package main

// Xapian - C++ search engine library
// Homepage: https://xapian.org/

import (
	"fmt"
	
	"os/exec"
)

func installXapian() {
	// Método 1: Descargar y extraer .tar.gz
	xapian_tar_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-core-1.4.26.tar.xz"
	xapian_cmd_tar := exec.Command("curl", "-L", xapian_tar_url, "-o", "package.tar.gz")
	err := xapian_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xapian_zip_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-core-1.4.26.tar.xz"
	xapian_cmd_zip := exec.Command("curl", "-L", xapian_zip_url, "-o", "package.zip")
	err = xapian_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xapian_bin_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-core-1.4.26.tar.xz"
	xapian_cmd_bin := exec.Command("curl", "-L", xapian_bin_url, "-o", "binary.bin")
	err = xapian_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xapian_src_url := "https://oligarchy.co.uk/xapian/1.4.26/xapian-core-1.4.26.tar.xz"
	xapian_cmd_src := exec.Command("curl", "-L", xapian_src_url, "-o", "source.tar.gz")
	err = xapian_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xapian_cmd_direct := exec.Command("./binary")
	err = xapian_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: util-linux")
exec.Command("latte", "install", "util-linux")
}
