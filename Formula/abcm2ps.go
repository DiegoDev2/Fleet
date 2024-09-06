package main

// Abcm2ps - ABC music notation software
// Homepage: http://moinejf.free.fr

import (
	"fmt"
	
	"os/exec"
)

func installAbcm2ps() {
	// Método 1: Descargar y extraer .tar.gz
	abcm2ps_tar_url := "https://github.com/leesavide/abcm2ps/archive/refs/tags/v8.14.15.tar.gz"
	abcm2ps_cmd_tar := exec.Command("curl", "-L", abcm2ps_tar_url, "-o", "package.tar.gz")
	err := abcm2ps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abcm2ps_zip_url := "https://github.com/leesavide/abcm2ps/archive/refs/tags/v8.14.15.zip"
	abcm2ps_cmd_zip := exec.Command("curl", "-L", abcm2ps_zip_url, "-o", "package.zip")
	err = abcm2ps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abcm2ps_bin_url := "https://github.com/leesavide/abcm2ps/archive/refs/tags/v8.14.15.bin"
	abcm2ps_cmd_bin := exec.Command("curl", "-L", abcm2ps_bin_url, "-o", "binary.bin")
	err = abcm2ps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abcm2ps_src_url := "https://github.com/leesavide/abcm2ps/archive/refs/tags/v8.14.15.src.tar.gz"
	abcm2ps_cmd_src := exec.Command("curl", "-L", abcm2ps_src_url, "-o", "source.tar.gz")
	err = abcm2ps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abcm2ps_cmd_direct := exec.Command("./binary")
	err = abcm2ps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
