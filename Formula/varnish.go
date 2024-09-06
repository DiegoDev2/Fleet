package main

// Varnish - High-performance HTTP accelerator
// Homepage: https://www.varnish-cache.org/

import (
	"fmt"
	
	"os/exec"
)

func installVarnish() {
	// Método 1: Descargar y extraer .tar.gz
	varnish_tar_url := "https://varnish-cache.org/_downloads/varnish-7.5.0.tgz"
	varnish_cmd_tar := exec.Command("curl", "-L", varnish_tar_url, "-o", "package.tar.gz")
	err := varnish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	varnish_zip_url := "https://varnish-cache.org/_downloads/varnish-7.5.0.tgz"
	varnish_cmd_zip := exec.Command("curl", "-L", varnish_zip_url, "-o", "package.zip")
	err = varnish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	varnish_bin_url := "https://varnish-cache.org/_downloads/varnish-7.5.0.tgz"
	varnish_cmd_bin := exec.Command("curl", "-L", varnish_bin_url, "-o", "binary.bin")
	err = varnish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	varnish_src_url := "https://varnish-cache.org/_downloads/varnish-7.5.0.tgz"
	varnish_cmd_src := exec.Command("curl", "-L", varnish_src_url, "-o", "source.tar.gz")
	err = varnish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	varnish_cmd_direct := exec.Command("./binary")
	err = varnish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docutils")
	exec.Command("latte", "install", "docutils").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
}
