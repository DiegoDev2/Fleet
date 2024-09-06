package main

// Hbase - Hadoop database: a distributed, scalable, big data store
// Homepage: https://hbase.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installHbase() {
	// Método 1: Descargar y extraer .tar.gz
	hbase_tar_url := "https://www.apache.org/dyn/closer.lua?path=hbase/2.6.0/hbase-2.6.0-bin.tar.gz"
	hbase_cmd_tar := exec.Command("curl", "-L", hbase_tar_url, "-o", "package.tar.gz")
	err := hbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hbase_zip_url := "https://www.apache.org/dyn/closer.lua?path=hbase/2.6.0/hbase-2.6.0-bin.zip"
	hbase_cmd_zip := exec.Command("curl", "-L", hbase_zip_url, "-o", "package.zip")
	err = hbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hbase_bin_url := "https://www.apache.org/dyn/closer.lua?path=hbase/2.6.0/hbase-2.6.0-bin.bin"
	hbase_cmd_bin := exec.Command("curl", "-L", hbase_bin_url, "-o", "binary.bin")
	err = hbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hbase_src_url := "https://www.apache.org/dyn/closer.lua?path=hbase/2.6.0/hbase-2.6.0-bin.src.tar.gz"
	hbase_cmd_src := exec.Command("curl", "-L", hbase_src_url, "-o", "source.tar.gz")
	err = hbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hbase_cmd_direct := exec.Command("./binary")
	err = hbase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
	exec.Command("latte", "install", "ant").Run()
	fmt.Println("Instalando dependencia: lzo")
	exec.Command("latte", "install", "lzo").Run()
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
