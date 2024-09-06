package main

// Jruby - Ruby implementation in pure Java
// Homepage: https://www.jruby.org/

import (
	"fmt"
	
	"os/exec"
)

func installJruby() {
	// Método 1: Descargar y extraer .tar.gz
	jruby_tar_url := "https://search.maven.org/remotecontent?filepath=org/jruby/jruby-dist/9.4.8.0/jruby-dist-9.4.8.0-bin.tar.gz"
	jruby_cmd_tar := exec.Command("curl", "-L", jruby_tar_url, "-o", "package.tar.gz")
	err := jruby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jruby_zip_url := "https://search.maven.org/remotecontent?filepath=org/jruby/jruby-dist/9.4.8.0/jruby-dist-9.4.8.0-bin.zip"
	jruby_cmd_zip := exec.Command("curl", "-L", jruby_zip_url, "-o", "package.zip")
	err = jruby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jruby_bin_url := "https://search.maven.org/remotecontent?filepath=org/jruby/jruby-dist/9.4.8.0/jruby-dist-9.4.8.0-bin.bin"
	jruby_cmd_bin := exec.Command("curl", "-L", jruby_bin_url, "-o", "binary.bin")
	err = jruby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jruby_src_url := "https://search.maven.org/remotecontent?filepath=org/jruby/jruby-dist/9.4.8.0/jruby-dist-9.4.8.0-bin.src.tar.gz"
	jruby_cmd_src := exec.Command("curl", "-L", jruby_src_url, "-o", "source.tar.gz")
	err = jruby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jruby_cmd_direct := exec.Command("./binary")
	err = jruby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
