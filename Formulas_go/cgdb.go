package main

// Cgdb - Curses-based interface to the GNU Debugger
// Homepage: https://cgdb.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCgdb() {
	// Método 1: Descargar y extraer .tar.gz
	cgdb_tar_url := "https://cgdb.me/files/cgdb-0.8.0.tar.gz"
	cgdb_cmd_tar := exec.Command("curl", "-L", cgdb_tar_url, "-o", "package.tar.gz")
	err := cgdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cgdb_zip_url := "https://cgdb.me/files/cgdb-0.8.0.zip"
	cgdb_cmd_zip := exec.Command("curl", "-L", cgdb_zip_url, "-o", "package.zip")
	err = cgdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cgdb_bin_url := "https://cgdb.me/files/cgdb-0.8.0.bin"
	cgdb_cmd_bin := exec.Command("curl", "-L", cgdb_bin_url, "-o", "binary.bin")
	err = cgdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cgdb_src_url := "https://cgdb.me/files/cgdb-0.8.0.src.tar.gz"
	cgdb_cmd_src := exec.Command("curl", "-L", cgdb_src_url, "-o", "source.tar.gz")
	err = cgdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cgdb_cmd_direct := exec.Command("./binary")
	err = cgdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
