package main

// Sbuild - Scala-based build system
// Homepage: http://sbuild.org/

import (
	"fmt"
	
	"os/exec"
)

func installSbuild() {
	// Método 1: Descargar y extraer .tar.gz
	sbuild_tar_url := "http://sbuild.org/uploads/sbuild/0.7.7/sbuild-0.7.7-dist.zip"
	sbuild_cmd_tar := exec.Command("curl", "-L", sbuild_tar_url, "-o", "package.tar.gz")
	err := sbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbuild_zip_url := "http://sbuild.org/uploads/sbuild/0.7.7/sbuild-0.7.7-dist.zip"
	sbuild_cmd_zip := exec.Command("curl", "-L", sbuild_zip_url, "-o", "package.zip")
	err = sbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbuild_bin_url := "http://sbuild.org/uploads/sbuild/0.7.7/sbuild-0.7.7-dist.zip"
	sbuild_cmd_bin := exec.Command("curl", "-L", sbuild_bin_url, "-o", "binary.bin")
	err = sbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbuild_src_url := "http://sbuild.org/uploads/sbuild/0.7.7/sbuild-0.7.7-dist.zip"
	sbuild_cmd_src := exec.Command("curl", "-L", sbuild_src_url, "-o", "source.tar.gz")
	err = sbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbuild_cmd_direct := exec.Command("./binary")
	err = sbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
