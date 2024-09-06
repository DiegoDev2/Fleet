package main

// Gearman - Application framework to farm out work to other machines or processes
// Homepage: https://gearman.org/

import (
	"fmt"
	
	"os/exec"
)

func installGearman() {
	// Método 1: Descargar y extraer .tar.gz
	gearman_tar_url := "https://github.com/gearman/gearmand/releases/download/1.1.21/gearmand-1.1.21.tar.gz"
	gearman_cmd_tar := exec.Command("curl", "-L", gearman_tar_url, "-o", "package.tar.gz")
	err := gearman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gearman_zip_url := "https://github.com/gearman/gearmand/releases/download/1.1.21/gearmand-1.1.21.zip"
	gearman_cmd_zip := exec.Command("curl", "-L", gearman_zip_url, "-o", "package.zip")
	err = gearman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gearman_bin_url := "https://github.com/gearman/gearmand/releases/download/1.1.21/gearmand-1.1.21.bin"
	gearman_cmd_bin := exec.Command("curl", "-L", gearman_bin_url, "-o", "binary.bin")
	err = gearman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gearman_src_url := "https://github.com/gearman/gearmand/releases/download/1.1.21/gearmand-1.1.21.src.tar.gz"
	gearman_cmd_src := exec.Command("curl", "-L", gearman_src_url, "-o", "source.tar.gz")
	err = gearman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gearman_cmd_direct := exec.Command("./binary")
	err = gearman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libmemcached")
	exec.Command("latte", "install", "libmemcached").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
