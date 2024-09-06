package main

// MavenShell - Shell for Maven
// Homepage: https://github.com/jdillon/mvnsh

import (
	"fmt"
	
	"os/exec"
)

func installMavenShell() {
	// Método 1: Descargar y extraer .tar.gz
	mavenshell_tar_url := "https://search.maven.org/remotecontent?filepath=org/sonatype/maven/shell/dist/mvnsh-assembly/1.1.0/mvnsh-assembly-1.1.0-bin.tar.gz"
	mavenshell_cmd_tar := exec.Command("curl", "-L", mavenshell_tar_url, "-o", "package.tar.gz")
	err := mavenshell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mavenshell_zip_url := "https://search.maven.org/remotecontent?filepath=org/sonatype/maven/shell/dist/mvnsh-assembly/1.1.0/mvnsh-assembly-1.1.0-bin.zip"
	mavenshell_cmd_zip := exec.Command("curl", "-L", mavenshell_zip_url, "-o", "package.zip")
	err = mavenshell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mavenshell_bin_url := "https://search.maven.org/remotecontent?filepath=org/sonatype/maven/shell/dist/mvnsh-assembly/1.1.0/mvnsh-assembly-1.1.0-bin.bin"
	mavenshell_cmd_bin := exec.Command("curl", "-L", mavenshell_bin_url, "-o", "binary.bin")
	err = mavenshell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mavenshell_src_url := "https://search.maven.org/remotecontent?filepath=org/sonatype/maven/shell/dist/mvnsh-assembly/1.1.0/mvnsh-assembly-1.1.0-bin.src.tar.gz"
	mavenshell_cmd_src := exec.Command("curl", "-L", mavenshell_src_url, "-o", "source.tar.gz")
	err = mavenshell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mavenshell_cmd_direct := exec.Command("./binary")
	err = mavenshell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
