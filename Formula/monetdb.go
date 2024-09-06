package main

// Monetdb - Column-store database
// Homepage: https://www.monetdb.org/

import (
	"fmt"
	
	"os/exec"
)

func installMonetdb() {
	// Método 1: Descargar y extraer .tar.gz
	monetdb_tar_url := "https://www.monetdb.org/downloads/sources/Aug2024/MonetDB-11.51.3.tar.xz"
	monetdb_cmd_tar := exec.Command("curl", "-L", monetdb_tar_url, "-o", "package.tar.gz")
	err := monetdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monetdb_zip_url := "https://www.monetdb.org/downloads/sources/Aug2024/MonetDB-11.51.3.tar.xz"
	monetdb_cmd_zip := exec.Command("curl", "-L", monetdb_zip_url, "-o", "package.zip")
	err = monetdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monetdb_bin_url := "https://www.monetdb.org/downloads/sources/Aug2024/MonetDB-11.51.3.tar.xz"
	monetdb_cmd_bin := exec.Command("curl", "-L", monetdb_bin_url, "-o", "binary.bin")
	err = monetdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monetdb_src_url := "https://www.monetdb.org/downloads/sources/Aug2024/MonetDB-11.51.3.tar.xz"
	monetdb_cmd_src := exec.Command("curl", "-L", monetdb_src_url, "-o", "source.tar.gz")
	err = monetdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monetdb_cmd_direct := exec.Command("./binary")
	err = monetdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
