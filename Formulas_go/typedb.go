package main

// Typedb - Strongly-typed database with a rich and logical type system
// Homepage: https://vaticle.com/

import (
	"fmt"
	
	"os/exec"
)

func installTypedb() {
	// Método 1: Descargar y extraer .tar.gz
	typedb_tar_url := "https://github.com/vaticle/typedb/releases/download/2.23.0/typedb-all-mac-2.23.0.zip"
	typedb_cmd_tar := exec.Command("curl", "-L", typedb_tar_url, "-o", "package.tar.gz")
	err := typedb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typedb_zip_url := "https://github.com/vaticle/typedb/releases/download/2.23.0/typedb-all-mac-2.23.0.zip"
	typedb_cmd_zip := exec.Command("curl", "-L", typedb_zip_url, "-o", "package.zip")
	err = typedb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typedb_bin_url := "https://github.com/vaticle/typedb/releases/download/2.23.0/typedb-all-mac-2.23.0.zip"
	typedb_cmd_bin := exec.Command("curl", "-L", typedb_bin_url, "-o", "binary.bin")
	err = typedb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typedb_src_url := "https://github.com/vaticle/typedb/releases/download/2.23.0/typedb-all-mac-2.23.0.zip"
	typedb_cmd_src := exec.Command("curl", "-L", typedb_src_url, "-o", "source.tar.gz")
	err = typedb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typedb_cmd_direct := exec.Command("./binary")
	err = typedb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
