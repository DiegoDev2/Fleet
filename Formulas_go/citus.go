package main

// Citus - PostgreSQL-based distributed RDBMS
// Homepage: https://www.citusdata.com

import (
	"fmt"
	
	"os/exec"
)

func installCitus() {
	// Método 1: Descargar y extraer .tar.gz
	citus_tar_url := "https://github.com/citusdata/citus/archive/refs/tags/v12.1.4.tar.gz"
	citus_cmd_tar := exec.Command("curl", "-L", citus_tar_url, "-o", "package.tar.gz")
	err := citus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	citus_zip_url := "https://github.com/citusdata/citus/archive/refs/tags/v12.1.4.zip"
	citus_cmd_zip := exec.Command("curl", "-L", citus_zip_url, "-o", "package.zip")
	err = citus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	citus_bin_url := "https://github.com/citusdata/citus/archive/refs/tags/v12.1.4.bin"
	citus_cmd_bin := exec.Command("curl", "-L", citus_bin_url, "-o", "binary.bin")
	err = citus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	citus_src_url := "https://github.com/citusdata/citus/archive/refs/tags/v12.1.4.src.tar.gz"
	citus_cmd_src := exec.Command("curl", "-L", citus_src_url, "-o", "source.tar.gz")
	err = citus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	citus_cmd_direct := exec.Command("./binary")
	err = citus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: postgresql@14")
exec.Command("latte", "install", "postgresql@14")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
