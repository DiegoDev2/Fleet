package main

// GradleProfiler - Profiling and benchmarking tool for Gradle builds
// Homepage: https://github.com/gradle/gradle-profiler/

import (
	"fmt"
	
	"os/exec"
)

func installGradleProfiler() {
	// Método 1: Descargar y extraer .tar.gz
	gradleprofiler_tar_url := "https://search.maven.org/remotecontent?filepath=org/gradle/profiler/gradle-profiler/0.20.0/gradle-profiler-0.20.0.zip"
	gradleprofiler_cmd_tar := exec.Command("curl", "-L", gradleprofiler_tar_url, "-o", "package.tar.gz")
	err := gradleprofiler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gradleprofiler_zip_url := "https://search.maven.org/remotecontent?filepath=org/gradle/profiler/gradle-profiler/0.20.0/gradle-profiler-0.20.0.zip"
	gradleprofiler_cmd_zip := exec.Command("curl", "-L", gradleprofiler_zip_url, "-o", "package.zip")
	err = gradleprofiler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gradleprofiler_bin_url := "https://search.maven.org/remotecontent?filepath=org/gradle/profiler/gradle-profiler/0.20.0/gradle-profiler-0.20.0.zip"
	gradleprofiler_cmd_bin := exec.Command("curl", "-L", gradleprofiler_bin_url, "-o", "binary.bin")
	err = gradleprofiler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gradleprofiler_src_url := "https://search.maven.org/remotecontent?filepath=org/gradle/profiler/gradle-profiler/0.20.0/gradle-profiler-0.20.0.zip"
	gradleprofiler_cmd_src := exec.Command("curl", "-L", gradleprofiler_src_url, "-o", "source.tar.gz")
	err = gradleprofiler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gradleprofiler_cmd_direct := exec.Command("./binary")
	err = gradleprofiler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
