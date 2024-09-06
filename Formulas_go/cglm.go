package main

// Cglm - Optimized OpenGL/Graphics Math (glm) for C
// Homepage: https://github.com/recp/cglm

import (
	"fmt"
	
	"os/exec"
)

func installCglm() {
	// Método 1: Descargar y extraer .tar.gz
	cglm_tar_url := "https://github.com/recp/cglm/archive/refs/tags/v0.9.4.tar.gz"
	cglm_cmd_tar := exec.Command("curl", "-L", cglm_tar_url, "-o", "package.tar.gz")
	err := cglm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cglm_zip_url := "https://github.com/recp/cglm/archive/refs/tags/v0.9.4.zip"
	cglm_cmd_zip := exec.Command("curl", "-L", cglm_zip_url, "-o", "package.zip")
	err = cglm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cglm_bin_url := "https://github.com/recp/cglm/archive/refs/tags/v0.9.4.bin"
	cglm_cmd_bin := exec.Command("curl", "-L", cglm_bin_url, "-o", "binary.bin")
	err = cglm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cglm_src_url := "https://github.com/recp/cglm/archive/refs/tags/v0.9.4.src.tar.gz"
	cglm_cmd_src := exec.Command("curl", "-L", cglm_src_url, "-o", "source.tar.gz")
	err = cglm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cglm_cmd_direct := exec.Command("./binary")
	err = cglm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
