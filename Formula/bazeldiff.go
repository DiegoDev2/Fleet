package main

// BazelDiff - Performs Bazel Target Diffing between two revisions in Git
// Homepage: https://github.com/Tinder/bazel-diff/

import (
	"fmt"
	
	"os/exec"
)

func installBazelDiff() {
	// Método 1: Descargar y extraer .tar.gz
	bazeldiff_tar_url := "https://github.com/Tinder/bazel-diff/releases/download/7.1.1/bazel-diff_deploy.jar"
	bazeldiff_cmd_tar := exec.Command("curl", "-L", bazeldiff_tar_url, "-o", "package.tar.gz")
	err := bazeldiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bazeldiff_zip_url := "https://github.com/Tinder/bazel-diff/releases/download/7.1.1/bazel-diff_deploy.jar"
	bazeldiff_cmd_zip := exec.Command("curl", "-L", bazeldiff_zip_url, "-o", "package.zip")
	err = bazeldiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bazeldiff_bin_url := "https://github.com/Tinder/bazel-diff/releases/download/7.1.1/bazel-diff_deploy.jar"
	bazeldiff_cmd_bin := exec.Command("curl", "-L", bazeldiff_bin_url, "-o", "binary.bin")
	err = bazeldiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bazeldiff_src_url := "https://github.com/Tinder/bazel-diff/releases/download/7.1.1/bazel-diff_deploy.jar"
	bazeldiff_cmd_src := exec.Command("curl", "-L", bazeldiff_src_url, "-o", "source.tar.gz")
	err = bazeldiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bazeldiff_cmd_direct := exec.Command("./binary")
	err = bazeldiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bazel")
	exec.Command("latte", "install", "bazel").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
