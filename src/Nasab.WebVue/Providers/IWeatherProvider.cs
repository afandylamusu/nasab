using System.Collections.Generic;
using Nasab.WebVue.Models;

namespace Nasab.WebVue.Providers
{
    public interface IWeatherProvider
    {
        List<WeatherForecast> GetForecasts();
    }
}
