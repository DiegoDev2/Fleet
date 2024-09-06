package main

// Psqlodbc - Official PostgreSQL ODBC driver
// Homepage: https://odbc.postgresql.org

import (
	"fmt"
	
	"os/exec"
)

func installPsqlodbc() {
	// Método 1: Descargar y extraer .tar.gz
	psqlodbc_tar_url := "https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-16_00_0005.tar.gz"
	psqlodbc_cmd_tar := exec.Command("curl", "-L", psqlodbc_tar_url, "-o", "package.tar.gz")
	err := psqlodbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psqlodbc_zip_url := "https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-16_00_0005.zip"
	psqlodbc_cmd_zip := exec.Command("curl", "-L", psqlodbc_zip_url, "-o", "package.zip")
	err = psqlodbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psqlodbc_bin_url := "https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-16_00_0005.bin"
	psqlodbc_cmd_bin := exec.Command("curl", "-L", psqlodbc_bin_url, "-o", "binary.bin")
	err = psqlodbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psqlodbc_src_url := "https://github.com/postgresql-interfaces/psqlodbc/archive/refs/tags/REL-16_00_0005.src.tar.gz"
	psqlodbc_cmd_src := exec.Command("curl", "-L", psqlodbc_src_url, "-o", "source.tar.gz")
	err = psqlodbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psqlodbc_cmd_direct := exec.Command("./binary")
	err = psqlodbc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
}
