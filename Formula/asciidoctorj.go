package main

// Asciidoctorj - Java wrapper and bindings for Asciidoctor
// Homepage: https://github.com/asciidoctor/asciidoctorj

import (
	"fmt"
	
	"os/exec"
)

func installAsciidoctorj() {
	// Método 1: Descargar y extraer .tar.gz
	asciidoctorj_tar_url := "https://search.maven.org/remotecontent?filepath=org/asciidoctor/asciidoctorj/3.0.0/asciidoctorj-3.0.0-bin.zip"
	asciidoctorj_cmd_tar := exec.Command("curl", "-L", asciidoctorj_tar_url, "-o", "package.tar.gz")
	err := asciidoctorj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciidoctorj_zip_url := "https://search.maven.org/remotecontent?filepath=org/asciidoctor/asciidoctorj/3.0.0/asciidoctorj-3.0.0-bin.zip"
	asciidoctorj_cmd_zip := exec.Command("curl", "-L", asciidoctorj_zip_url, "-o", "package.zip")
	err = asciidoctorj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciidoctorj_bin_url := "https://search.maven.org/remotecontent?filepath=org/asciidoctor/asciidoctorj/3.0.0/asciidoctorj-3.0.0-bin.zip"
	asciidoctorj_cmd_bin := exec.Command("curl", "-L", asciidoctorj_bin_url, "-o", "binary.bin")
	err = asciidoctorj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciidoctorj_src_url := "https://search.maven.org/remotecontent?filepath=org/asciidoctor/asciidoctorj/3.0.0/asciidoctorj-3.0.0-bin.zip"
	asciidoctorj_cmd_src := exec.Command("curl", "-L", asciidoctorj_src_url, "-o", "source.tar.gz")
	err = asciidoctorj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciidoctorj_cmd_direct := exec.Command("./binary")
	err = asciidoctorj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
