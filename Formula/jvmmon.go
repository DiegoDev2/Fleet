package main

// JvmMon - Console-based JVM monitoring
// Homepage: https://github.com/ajermakovics/jvm-mon

import (
	"fmt"
	
	"os/exec"
)

func installJvmMon() {
	// Método 1: Descargar y extraer .tar.gz
	jvmmon_tar_url := "https://github.com/ajermakovics/jvm-mon/releases/download/0.3/jvm-mon-0.3.tar.gz"
	jvmmon_cmd_tar := exec.Command("curl", "-L", jvmmon_tar_url, "-o", "package.tar.gz")
	err := jvmmon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jvmmon_zip_url := "https://github.com/ajermakovics/jvm-mon/releases/download/0.3/jvm-mon-0.3.zip"
	jvmmon_cmd_zip := exec.Command("curl", "-L", jvmmon_zip_url, "-o", "package.zip")
	err = jvmmon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jvmmon_bin_url := "https://github.com/ajermakovics/jvm-mon/releases/download/0.3/jvm-mon-0.3.bin"
	jvmmon_cmd_bin := exec.Command("curl", "-L", jvmmon_bin_url, "-o", "binary.bin")
	err = jvmmon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jvmmon_src_url := "https://github.com/ajermakovics/jvm-mon/releases/download/0.3/jvm-mon-0.3.src.tar.gz"
	jvmmon_cmd_src := exec.Command("curl", "-L", jvmmon_src_url, "-o", "source.tar.gz")
	err = jvmmon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jvmmon_cmd_direct := exec.Command("./binary")
	err = jvmmon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@8")
	exec.Command("latte", "install", "openjdk@8").Run()
}
