package main

// Gitbucket - Git platform powered by Scala offering
// Homepage: https://github.com/gitbucket/gitbucket

import (
	"fmt"
	
	"os/exec"
)

func installGitbucket() {
	// Método 1: Descargar y extraer .tar.gz
	gitbucket_tar_url := "https://github.com/gitbucket/gitbucket/releases/download/4.41.0/gitbucket.war"
	gitbucket_cmd_tar := exec.Command("curl", "-L", gitbucket_tar_url, "-o", "package.tar.gz")
	err := gitbucket_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitbucket_zip_url := "https://github.com/gitbucket/gitbucket/releases/download/4.41.0/gitbucket.war"
	gitbucket_cmd_zip := exec.Command("curl", "-L", gitbucket_zip_url, "-o", "package.zip")
	err = gitbucket_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitbucket_bin_url := "https://github.com/gitbucket/gitbucket/releases/download/4.41.0/gitbucket.war"
	gitbucket_cmd_bin := exec.Command("curl", "-L", gitbucket_bin_url, "-o", "binary.bin")
	err = gitbucket_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitbucket_src_url := "https://github.com/gitbucket/gitbucket/releases/download/4.41.0/gitbucket.war"
	gitbucket_cmd_src := exec.Command("curl", "-L", gitbucket_src_url, "-o", "source.tar.gz")
	err = gitbucket_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitbucket_cmd_direct := exec.Command("./binary")
	err = gitbucket_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sbt")
exec.Command("latte", "install", "sbt")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
