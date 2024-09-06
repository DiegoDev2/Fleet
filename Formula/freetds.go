package main

// Freetds - Libraries to talk to Microsoft SQL Server and Sybase databases
// Homepage: https://www.freetds.org/

import (
	"fmt"
	
	"os/exec"
)

func installFreetds() {
	// Método 1: Descargar y extraer .tar.gz
	freetds_tar_url := "https://www.freetds.org/files/stable/freetds-1.4.22.tar.bz2"
	freetds_cmd_tar := exec.Command("curl", "-L", freetds_tar_url, "-o", "package.tar.gz")
	err := freetds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freetds_zip_url := "https://www.freetds.org/files/stable/freetds-1.4.22.tar.bz2"
	freetds_cmd_zip := exec.Command("curl", "-L", freetds_zip_url, "-o", "package.zip")
	err = freetds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freetds_bin_url := "https://www.freetds.org/files/stable/freetds-1.4.22.tar.bz2"
	freetds_cmd_bin := exec.Command("curl", "-L", freetds_bin_url, "-o", "binary.bin")
	err = freetds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freetds_src_url := "https://www.freetds.org/files/stable/freetds-1.4.22.tar.bz2"
	freetds_cmd_src := exec.Command("curl", "-L", freetds_src_url, "-o", "source.tar.gz")
	err = freetds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freetds_cmd_direct := exec.Command("./binary")
	err = freetds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
