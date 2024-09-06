package main

// Tsung - Load testing for HTTP, PostgreSQL, Jabber, and others
// Homepage: http://tsung.erlang-projects.org/

import (
	"fmt"
	
	"os/exec"
)

func installTsung() {
	// Método 1: Descargar y extraer .tar.gz
	tsung_tar_url := "http://tsung.erlang-projects.org/dist/tsung-1.8.0.tar.gz"
	tsung_cmd_tar := exec.Command("curl", "-L", tsung_tar_url, "-o", "package.tar.gz")
	err := tsung_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tsung_zip_url := "http://tsung.erlang-projects.org/dist/tsung-1.8.0.zip"
	tsung_cmd_zip := exec.Command("curl", "-L", tsung_zip_url, "-o", "package.zip")
	err = tsung_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tsung_bin_url := "http://tsung.erlang-projects.org/dist/tsung-1.8.0.bin"
	tsung_cmd_bin := exec.Command("curl", "-L", tsung_bin_url, "-o", "binary.bin")
	err = tsung_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tsung_src_url := "http://tsung.erlang-projects.org/dist/tsung-1.8.0.src.tar.gz"
	tsung_cmd_src := exec.Command("curl", "-L", tsung_src_url, "-o", "source.tar.gz")
	err = tsung_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tsung_cmd_direct := exec.Command("./binary")
	err = tsung_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
exec.Command("latte", "install", "erlang")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
}
